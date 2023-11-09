{
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
  };
  outputs = { nixpkgs, ... }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        buildInputs = with pkgs; [
          glib
          graphene
          gdk-pixbuf
          gtk4
          vulkan-headers
        ];

        nativeBuildInputs = with pkgs; [
          gobject-introspection
          pkg-config
          go
        ];

        CGO_ENABLED = "1";
        ASSUME_NO_MOVING_GC_UNSAFE_RISK_IT_WITH = "go1.21";
      };
    };
}
