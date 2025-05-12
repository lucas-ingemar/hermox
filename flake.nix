{
  description = "hermox";

  inputs = { nixpkgs.url = "github:nixos/nixpkgs"; };

  outputs = { self, nixpkgs }:
    with nixpkgs.legacyPackages.x86_64-linux; {
      devShell.x86_64-linux =
        mkShell { buildInputs = [ rustc cargo rust-analyzer rustfmt ]; };
    };
}
