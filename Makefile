GOINSTALL = go install
BUILDFLAGS = -a
GITHUB_ACCOUNT = github.com/theeddieh
REPO_VIVEC = vivec
PACKAGE_SERVER = ${GITHUB_ACCOUNT}/${REPO_VIVEC}/server

.DEAFULT_GOAL: install

.PHONY: install
install:
	${GOINSTALL} ${PACKAGE_SERVER}
