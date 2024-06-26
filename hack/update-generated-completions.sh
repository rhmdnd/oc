#!/usr/bin/env bash

# This script sets up a go workspace locally and generates shell auto-completion scripts.

source "$(dirname "${BASH_SOURCE}")/lib/init.sh"

function os::build::gen-completions() {
  local dest="$1"
  local shell="$2"
  local skipprefix="${3:-}"

  # We do this in a tmpdir in case the dest has other non-autogenned files
  # We don't want to include them in the list of gen'd files
  local tmpdir="${OS_ROOT}/_tmp/gen_comp"
  mkdir -p "${tmpdir}"
  # generate the new files
  ${oc} completion ${shell} > $tmpdir/oc
  # create the list of generated files
  ls "${tmpdir}" | LC_ALL=C sort > "${tmpdir}/.files_generated"

  if [[ -e "${dest}/.files_generated" ]]; then
    # remove all old generated file from the destination
    while read file; do
      if [[ -e "${tmpdir}/${file}" && -n "${skipprefix}" ]]; then
        local original generated
        original=$(grep -v "^${skipprefix}" "${dest}/${file}") || :
        generated=$(grep -v "^${skipprefix}" "${tmpdir}/${file}") || :
        if [[ "${original}" == "${generated}" ]]; then
          # overwrite generated with original.
          mv "${dest}/${file}" "${tmpdir}/${file}"
        fi
      else
        rm "${dest}/${file}" || true
      fi
    done <"${dest}/.files_generated"
  fi
  # put the new generated file into the destination
  find "${tmpdir}" -exec rsync -pt {} "${dest}" \; >/dev/null
  #cleanup
  rm -rf "${tmpdir}"

  echo "Assets generated in ${dest}"
}
readonly -f os::build::gen-completions

oc="${OS_ROOT}/oc"
platform="$(os::build::host_platform)"
if [[ "${platform}" != "linux/amd64" ]]; then
  os::log::warning "Generating completions on ${platform} may not be identical to running on linux/amd64 due to conditional compilation."
  oc="${OS_ROOT}/_output/bin/${platform//\//_}/oc"
fi

OUTPUT_DIR_ROOT="${1:-${OS_ROOT}/contrib/completions}"

mkdir -p "${OUTPUT_DIR_ROOT}/bash" || echo $? > /dev/null
mkdir -p "${OUTPUT_DIR_ROOT}/zsh" || echo $? > /dev/null

os::build::gen-completions "${OUTPUT_DIR_ROOT}/bash" "bash"
os::build::gen-completions "${OUTPUT_DIR_ROOT}/zsh" "zsh"
