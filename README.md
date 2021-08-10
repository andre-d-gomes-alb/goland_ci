# CI/CD with GitHub Actions in Go

### [Workflow](https://github.com/andre-d-gomes/goland_ci/blob/main/.github/workflows/cicd.yml):
  1. **Test syntax** by running a linting check with [golangci-lint](https://golangci-lint.run/). Lint is a static code analysis tool used to flag programming errors, bugs, stylistic errors and suspicious constructs.
  2. **Test functionality** by running automated **[unit tests](https://github.com/andre-d-gomes/goland_ci/blob/main/main_test.go)** on the entire program.
  3. **Checks** which **[version of the last release](https://github.com/andre-d-gomes/goland_ci/releases)** was executed and which **[version in the commit](https://github.com/andre-d-gomes/goland_ci/blob/main/VERSION)**.
  4. **Test build stability** by attempting to build the program for Linux and Windows.
      - **If the version was updated** on commit, the resulting build **binaries are saved in the release**.
  5. **If the version was updated** on commit, **creates a release of the new version** according to a [release template](https://github.com/andre-d-gomes/goland_ci/blob/main/.github/RELEASE-TEMPLATE.md).
 
