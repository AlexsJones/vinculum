#!/bin/sh
set -e
APP="vinculum"
REPO_NAME="AlexsJones/${APP}"
GITHUB_URL="https://github.com/$REPO_NAME/releases"
BIN_DIR="/usr/local/bin"
lastversion() {
  case $DOWNLOADER in
    curl)
      curl --silent "https://api.github.com/repos/$REPO_NAME/releases/latest" | # Get latest release from GitHub api
        grep '"tag_name":' |                                            # Get tag line
        sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
              ;;
            wget)
              wget -qO - "https://api.github.com/repos/$REPO_NAME/releases/latest" | # Get latest release from GitHub api
                grep '"tag_name":' |                                            # Get tag line
                sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
                              ;;
                            *)
                              fatal "Incorrect executable '$DOWNLOADER'"
                              ;;
                          esac

    # Abort if download command failed
    [ $? -eq 0 ] || fatal 'Download failed'
  }

# --- create temporary directory and cleanup when done ---
setup_tmp() {
  TMP_DIR=$(mktemp -d -t ${APP}-install.XXXXXXXXXX)
  TMP_BIN=${TMP_DIR}/${APP}.bin
  cleanup() {
    code=$?
    set +e
    trap - EXIT
    rm -rf "${TMP_DIR}"
    exit $code
  }
if [ -z "$DEBUG" ] 
  then 
    trap cleanup INT EXIT
  else 
    echo "ignoring deletion in debug mode of $TMP_DIR"
fi
}

# --- download from github url ---
download() {
  [ $# -eq 2 ] || fatal 'download needs exactly 2 arguments'

  case $DOWNLOADER in
    curl)
      curl -o $1 -sfL $2 --output /dev/stderr --write-out "%{http_code}"
      if [[ "$http_code" -ne 200 ]] ; then
        echo "$http_code"
        exit 1
      fi
      ;;
    wget)
      status_code=$(wget -qO $1 $2 --server-response 2>&1 | awk '/^  HTTP/{print $2}')
      if [[ "$status_code" -ne 200 ]] ; then
        echo "$status_code"
        exit 1
      fi
      ;;
    *)
      fatal "Incorrect executable '$DOWNLOADER'"
      ;;
  esac

    # Abort if download command failed
  }

# --- download binary from github url ---
download_binary() {
  BIN_URL=${GITHUB_URL}/download/${LATEST_VERSION}/${APP}_${OS}_${ARCH}
  echo "Downloading binary ${BIN_URL}"
  download ${TMP_BIN} ${BIN_URL}
}

setup_verify_os() {
  UNAME=$( command -v uname)
  case $( "${UNAME}" | tr '[:upper:]' '[:lower:]') in
    linux*)
      OS="linux"
      ;;
    darwin*)
      OS="darwin"
      ;;
    msys*|cygwin*|mingw*)
      # or possible 'bash on windows'
      OS="windows"
      ;;
    nt|win*)
      OS="windows"
      ;;
    *)
      printf 'unknown\n'
      ;;
  esac
  echo "${OS}"
}
setup_verify_arch() {
  if [ -z "$ARCH" ]; then
    ARCH=$(uname -m)
  else
    fatal "Unable to calculate architecture"
  fi
  case $ARCH in
    amd64)
      ARCH=amd64
      SUFFIX=
      ;;
    x86_64)
      ARCH=amd64
      SUFFIX=
      ;;
    arm64)
      ARCH=arm64
      SUFFIX=-${ARCH}
      ;;
    aarch64)
      ARCH=arm64
      SUFFIX=-${ARCH}
      ;;
    arm*)
      ARCH=arm
      SUFFIX=-${ARCH}hf
      ;;
    *)
      fatal "Unsupported architecture $ARCH"
  esac
  echo "${ARCH}"
}
# --- verify existence of network downloader executable ---
verify_downloader() {
  # Return failure if it doesn't exist or is no executable
  [ -x "$(which $1)" ] || return 1

    # Set verified executable as our downloader program and return success
    DOWNLOADER=$1
    return 0
  }

setup_binary() {

  chmod 755 ${TMP_BIN}
  echo "Installing ${APP} to ${BIN_DIR}/${APP}"
  $SUDO chown $(whoami) ${TMP_BIN}
  $SUDO mv -f "${TMP_BIN}" "${BIN_DIR}/${APP}"
}

setup_env() {
  SUDO=sudo
  if [ $(id -u) -eq 0 ]; then
    SUDO=
  fi
}
{
  setup_env
  verify_downloader curl || verify_downloader wget || fatal 'Can not find curl or wget for downloading files'
  LATEST_VERSION=$(lastversion "$DOWNLOADER")
  echo "Most recent remote version is $LATEST_VERSION"
  setup_verify_os
  echo "Local OS is $OS"
  setup_verify_arch
  echo "Local arch is $ARCH"
  setup_tmp
  download_binary "$LATEST_VERSION"
  setup_binary
}
