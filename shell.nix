# shell.nix
let
  pkgs = import <nixpkgs> {};
in
pkgs.mkShell {
  # buildInputs = [
  #   pkgs.hello
  # ];
  # nativeBuildInputs = [
  #   pkgs.go
  #   # dependencies you want available in your shell
  # ];
}
