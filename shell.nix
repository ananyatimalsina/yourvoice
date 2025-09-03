let pkgs = import <nixpkgs> { };
in pkgs.mkShell {
  packages = with pkgs; [ go nodejs biome rustywind ];
  shellHook = "export PATH=$PATH:$(go env GOPATH)/bin";
}
