{ pkgs, ... }:

{
  packages = [
    pkgs.git
    pkgs.go
    pkgs.gotools
    pkgs.go-task
    pkgs.golangci-lint
    pkgs.silver-searcher
  ];

  enterShell = ''
    git --version
    go version
    task --version
    golangci-lint --version
  '';

  # See full reference at https://devenv.sh/reference/options/
}
