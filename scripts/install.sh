#!/bin/sh
# install.sh - mdp installation script
# Usage: curl -fsSL https://raw.githubusercontent.com/sadiksaifi/mdp/main/scripts/install.sh | sh

set -e

# Configuration
REPO="sadiksaifi/mdp"
INSTALL_DIR="${HOME}/.local/bin"
DATA_DIR="${HOME}/.local/share/mdp"

# Colors (only if terminal supports them)
if [ -t 1 ]; then
    RED='\033[0;31m'
    GREEN='\033[0;32m'
    YELLOW='\033[1;33m'
    BLUE='\033[0;34m'
    NC='\033[0m' # No Color
else
    RED=''
    GREEN=''
    YELLOW=''
    BLUE=''
    NC=''
fi

info() {
    printf "${BLUE}==>${NC} %s\n" "$1"
}

success() {
    printf "${GREEN}==>${NC} %s\n" "$1"
}

warn() {
    printf "${YELLOW}Warning:${NC} %s\n" "$1"
}

error() {
    printf "${RED}Error:${NC} %s\n" "$1" >&2
    exit 1
}

# Detect OS
detect_os() {
    case "$(uname -s)" in
        Darwin)
            OS="darwin"
            ;;
        Linux)
            OS="linux"
            ;;
        MINGW*|MSYS*|CYGWIN*)
            error "Windows is not supported by this installer. Please download manually from GitHub."
            ;;
        *)
            error "Unsupported operating system: $(uname -s)"
            ;;
    esac
}

# Detect architecture
detect_arch() {
    case "$(uname -m)" in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        arm64|aarch64)
            ARCH="arm64"
            ;;
        armv7l|armv6l)
            error "ARM 32-bit is not supported. Please use ARM64 or AMD64."
            ;;
        *)
            error "Unsupported architecture: $(uname -m)"
            ;;
    esac
}

# Check for required tools
check_requirements() {
    if ! command -v curl >/dev/null 2>&1 && ! command -v wget >/dev/null 2>&1; then
        error "Neither curl nor wget found. Please install one of them."
    fi
}

# Get latest release version from GitHub
get_latest_version() {
    info "Fetching latest version..."

    local api_url="https://api.github.com/repos/${REPO}/releases/latest"

    if command -v curl >/dev/null 2>&1; then
        VERSION=$(curl -fsSL "$api_url" | grep -o '"tag_name": *"[^"]*"' | head -1 | sed 's/.*"tag_name": *"\([^"]*\)".*/\1/')
    else
        VERSION=$(wget -qO- "$api_url" | grep -o '"tag_name": *"[^"]*"' | head -1 | sed 's/.*"tag_name": *"\([^"]*\)".*/\1/')
    fi

    if [ -z "$VERSION" ]; then
        error "Failed to determine latest version. Please check your internet connection."
    fi

    info "Latest version: ${VERSION}"
}

# Download and install
download_and_install() {
    local url="https://github.com/${REPO}/releases/download/${VERSION}/mdp-${OS}-${ARCH}.tar.gz"
    local tmp_dir
    tmp_dir=$(mktemp -d)
    local archive="${tmp_dir}/mdp.tar.gz"

    info "Downloading mdp ${VERSION} for ${OS}/${ARCH}..."

    if command -v curl >/dev/null 2>&1; then
        if ! curl -fsSL "$url" -o "$archive"; then
            rm -rf "$tmp_dir"
            error "Download failed. The release for ${OS}/${ARCH} may not exist."
        fi
    else
        if ! wget -q "$url" -O "$archive"; then
            rm -rf "$tmp_dir"
            error "Download failed. The release for ${OS}/${ARCH} may not exist."
        fi
    fi

    info "Extracting..."
    if ! tar -xzf "$archive" -C "$tmp_dir" 2>/dev/null; then
        rm -rf "$tmp_dir"
        error "Extraction failed. The downloaded file may be corrupted."
    fi

    info "Installing to ${INSTALL_DIR}..."
    mkdir -p "$INSTALL_DIR"

    if [ -f "${INSTALL_DIR}/mdp" ]; then
        rm -f "${INSTALL_DIR}/mdp"
    fi

    mv "${tmp_dir}/mdp" "${INSTALL_DIR}/mdp"
    chmod +x "${INSTALL_DIR}/mdp"

    # Create marker file for installation method detection
    mkdir -p "$DATA_DIR"
    cat > "${DATA_DIR}/.curl-installed" << EOF
installed=$(date +%s)
version=${VERSION}
EOF

    # Cleanup
    rm -rf "$tmp_dir"

    success "mdp ${VERSION} installed successfully!"
}

# Configure PATH
configure_path() {
    # Check if already in PATH
    case ":$PATH:" in
        *":${INSTALL_DIR}:"*)
            return 0
            ;;
    esac

    info "Configuring PATH..."

    local shell_name
    shell_name=$(basename "$SHELL")
    local rc_file=""
    local path_line='export PATH="$HOME/.local/bin:$PATH"'

    case "$shell_name" in
        bash)
            # Check for .bash_profile first (macOS), then .bashrc (Linux)
            if [ -f "$HOME/.bash_profile" ]; then
                rc_file="$HOME/.bash_profile"
            elif [ -f "$HOME/.bashrc" ]; then
                rc_file="$HOME/.bashrc"
            else
                rc_file="$HOME/.bashrc"
            fi
            ;;
        zsh)
            if [ -f "$HOME/.zshrc" ]; then
                rc_file="$HOME/.zshrc"
            elif [ -f "${XDG_CONFIG_HOME:-$HOME/.config}/zsh/.zshrc" ]; then
                rc_file="${XDG_CONFIG_HOME:-$HOME/.config}/zsh/.zshrc"
            else
                rc_file="$HOME/.zshrc"
            fi
            ;;
        fish)
            rc_file="${XDG_CONFIG_HOME:-$HOME/.config}/fish/config.fish"
            path_line='set -gx PATH $HOME/.local/bin $PATH'
            ;;
        *)
            warn "Unknown shell: $shell_name"
            warn "Please add ${INSTALL_DIR} to your PATH manually:"
            warn "  export PATH=\"\$HOME/.local/bin:\$PATH\""
            return 0
            ;;
    esac

    # Check if path is already configured
    if [ -f "$rc_file" ] && grep -q '\.local/bin' "$rc_file" 2>/dev/null; then
        info "PATH already configured in ${rc_file}"
        return 0
    fi

    # Ensure rc file exists
    if [ ! -f "$rc_file" ]; then
        mkdir -p "$(dirname "$rc_file")"
        touch "$rc_file"
    fi

    # Add to rc file
    {
        echo ""
        echo "# Added by mdp installer"
        echo "$path_line"
    } >> "$rc_file"

    success "Added ${INSTALL_DIR} to PATH in ${rc_file}"
    echo ""
    warn "Run 'source ${rc_file}' or restart your terminal to use mdp"
}

# Verify installation
verify_installation() {
    if [ -x "${INSTALL_DIR}/mdp" ]; then
        local installed_version
        installed_version=$("${INSTALL_DIR}/mdp" --version 2>/dev/null | awk '{print $2}')
        success "Verified: mdp ${installed_version}"
    else
        error "Installation verification failed"
    fi
}

# Print post-install instructions
print_instructions() {
    echo ""
    echo "To get started, run:"
    echo ""
    echo "  mdp README.md"
    echo ""
    echo "For help, run:"
    echo ""
    echo "  mdp --help"
    echo ""
    echo "To upgrade later, run:"
    echo ""
    echo "  mdp upgrade"
    echo ""
}

# Main
main() {
    echo ""
    echo "mdp installer"
    echo "============="
    echo ""

    check_requirements
    detect_os
    detect_arch
    get_latest_version
    download_and_install
    configure_path
    verify_installation
    print_instructions
}

main
