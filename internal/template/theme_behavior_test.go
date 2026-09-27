package template

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dop251/goja"
)

const themeHarness = `
var __storage = {};
var __events = [];
var __buttons = [makeButton(), makeButton()];
var __media = {
    matches: false,
    listeners: {},
    addEventListener: function(type, listener) { this.listeners[type] = listener; }
};
function makeButton() {
    return {
        attributes: {},
        listeners: {},
        setAttribute: function(name, value) { this.attributes[name] = value; },
        addEventListener: function(type, listener) { this.listeners[type] = listener; }
    };
}
var document = {
    documentElement: { dataset: {} },
    querySelectorAll: function(selector) {
        if (selector !== '.theme-toggle-btn') throw new Error('unexpected selector: ' + selector);
        return __buttons;
    }
};
var localStorage = {
    getItem: function(key) {
        return Object.prototype.hasOwnProperty.call(__storage, key) ? __storage[key] : null;
    },
    setItem: function(key, value) { __storage[key] = String(value); },
    removeItem: function(key) { delete __storage[key]; }
};
var window = {
    matchMedia: function(query) {
        if (query !== '(prefers-color-scheme: dark)') throw new Error('unexpected query: ' + query);
        return __media;
    },
    dispatchEvent: function(event) { __events.push(event.type); }
};
function Event(type) { this.type = type; }
`

func TestGenerate_ThemeBehavior(t *testing.T) {
	t.Run("system preference is not persisted", func(t *testing.T) {
		vm := runGeneratedThemeScript(t, false, "")

		assertJSValue(t, vm, `document.documentElement.dataset.theme`, "light")
		assertJSValue(t, vm, `localStorage.getItem('mdp-theme')`, nil)
		assertJSValue(t, vm, `__buttons[0].attributes['aria-label']`, "Switch to dark mode")
		assertJSValue(t, vm, `__buttons[1].attributes['aria-label']`, "Switch to dark mode")
	})

	t.Run("dark system preference is applied without persistence", func(t *testing.T) {
		vm := runGeneratedThemeScript(t, true, "")

		assertJSValue(t, vm, `document.documentElement.dataset.theme`, "dark")
		assertJSValue(t, vm, `localStorage.getItem('mdp-theme')`, nil)
		assertJSValue(t, vm, `__buttons[0].attributes['aria-label']`, "Switch to light mode")
	})

	t.Run("mobile toggle persists and updates every control", func(t *testing.T) {
		vm := runGeneratedThemeScript(t, false, "")
		runJS(t, vm, `__buttons[1].listeners.click()`)

		assertJSValue(t, vm, `document.documentElement.dataset.theme`, "dark")
		assertJSValue(t, vm, `localStorage.getItem('mdp-theme')`, "dark")
		assertJSValue(t, vm, `__buttons[0].attributes['aria-label']`, "Switch to light mode")
		assertJSValue(t, vm, `__buttons[1].attributes['aria-label']`, "Switch to light mode")
		assertJSValue(t, vm, `__events[0]`, "mdp-theme-change")
	})

	t.Run("saved choice overrides system preference", func(t *testing.T) {
		vm := runGeneratedThemeScript(t, false, `__storage['mdp-theme'] = 'dark';`)

		assertJSValue(t, vm, `document.documentElement.dataset.theme`, "dark")
		assertJSValue(t, vm, `localStorage.getItem('mdp-theme')`, "dark")
	})

	t.Run("system changes apply until a manual choice", func(t *testing.T) {
		vm := runGeneratedThemeScript(t, false, "")
		runJS(t, vm, `__media.matches = true; __media.listeners.change()`)

		assertJSValue(t, vm, `document.documentElement.dataset.theme`, "dark")
		assertJSValue(t, vm, `localStorage.getItem('mdp-theme')`, nil)
	})

	t.Run("system changes do not override a manual choice", func(t *testing.T) {
		vm := runGeneratedThemeScript(t, false, "")
		runJS(t, vm, `__buttons[0].listeners.click()`)
		runJS(t, vm, `__media.matches = false; __media.listeners.change()`)

		assertJSValue(t, vm, `document.documentElement.dataset.theme`, "dark")
		assertJSValue(t, vm, `localStorage.getItem('mdp-theme')`, "dark")
	})
}

func runGeneratedThemeScript(t *testing.T, osDark bool, setup string) *goja.Runtime {
	t.Helper()

	output := Generate("Test", "<p>Content</p>")
	script := extractThemeScript(t, output)
	vm := goja.New()
	program := themeHarness + fmt.Sprintf("\n__media.matches = %t;\n", osDark) + setup + "\n" + script
	if _, err := vm.RunString(program); err != nil {
		t.Fatalf("execute generated theme script: %v", err)
	}
	return vm
}

func extractThemeScript(t *testing.T, output string) string {
	t.Helper()

	startMarker := `;(function() {
            var root = document.documentElement;`
	endMarker := `        })();`
	start := strings.LastIndex(output, startMarker)
	if start == -1 {
		t.Fatal("generated output does not contain the theme script")
	}
	end := strings.Index(output[start:], endMarker)
	if end == -1 {
		t.Fatal("generated theme script is not terminated")
	}
	return output[start : start+end+len(endMarker)]
}

func runJS(t *testing.T, vm *goja.Runtime, expression string) goja.Value {
	t.Helper()
	value, err := vm.RunString(expression)
	if err != nil {
		t.Fatalf("run JavaScript %q: %v", expression, err)
	}
	return value
}

func assertJSValue(t *testing.T, vm *goja.Runtime, expression string, want any) {
	t.Helper()
	got := runJS(t, vm, expression).Export()
	if got != want {
		t.Fatalf("%s = %#v, want %#v", expression, got, want)
	}
}
