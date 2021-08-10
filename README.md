# CI/CD with GitHub Actions in Go

### [Fluxo de trabalho](https://github.com/andre-d-gomes/goland_ci/blob/main/.github/workflows/cicd.yml):
  1. **Teste a sintaxe** executando uma verificação de linting com golangci-lint. Lint é uma ferramenta de análise de código estática usada para sinalizar erros de programação, bugs, erros estilísticos e construções suspeitas.
  2. **Teste a funcionalidade** pela execução de **[testes de unidade](https://github.com/andre-d-gomes/goland_ci/blob/main/main_test.go)** em todo o programa.
  3. **Verifica** qual é a **[versão da ultima release](https://github.com/andre-d-gomes/goland_ci/releases)** executada e qual a **[versão no commit](https://github.com/andre-d-gomes/goland_ci/blob/main/VERSION)**.
  4. **Teste a estabilidade de compilação** tentando compilar o programa para Linux and Windows.
      1. **Se a versão foi atualizada** no commit, **os binários** resultantes da compilação **são guardados na release**.
  5. **Se a versão foi atualizada** no commit, **cria uma release da nova versão** de acordo com um [template de release](https://github.com/andre-d-gomes/goland_ci/blob/main/.github/RELEASE-TEMPLATE.md).
 
