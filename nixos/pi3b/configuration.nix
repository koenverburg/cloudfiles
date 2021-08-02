{ config, pkgs, lib, ... }:
{
  imports = [
    <nixpkgs/nixos/modules/installer/cd-dvd/sd-image-aarch64.nix>

    # For nixpkgs cache
    <nixpkgs/nixos/modules/installer/cd-dvd/channel.nix>
  ];

  sdImage.compressImage = false;
  # NixOS wants to enable GRUB by default
  boot.loader.grub.enable = false;
  # Enables the generation of /boot/extlinux/extlinux.conf
  boot.loader.generic-extlinux-compatible.enable = true;
 
  # !!! Set to specific linux kernel version
  boot.kernelPackages = pkgs.linuxPackages_5_4;

  # !!! Needed for the virtual console to work on the RPi 3, as the default of 16M doesn't seem to be enough.
  # If X.org behaves weirdly (I only saw the cursor) then try increasing this to 256M.
  # On a Raspberry Pi 4 with 4 GB, you should either disable this parameter or increase to at least 64M if you want the USB ports to work.
  boot.kernelParams = ["cma=256M"];
    
  # File systems configuration for using the installer's partition layout
  fileSystems = {
    # Prior to 19.09, the boot partition was hosted on the smaller first partition
    # Starting with 19.09, the /boot folder is on the main bigger partition.
    # The following is to be used only with older images.
    /*
    "/boot" = {
      device = "/dev/disk/by-label/NIXOS_BOOT";
      fsType = "vfat";
    };
    */
    "/" = {
      device = "/dev/disk/by-label/NIXOS_SD";
      fsType = "ext4";
    };
  };

  # !!! Adding a swap file is optional, but strongly recommended!
  swapDevices = [ { device = "/swapfile"; size = 1024; } ];

  # systemPackages
  environment.systemPackages = with pkgs; [ 
    neovim curl wget bind kubectl helm iptables openvpn
    fish git python3 nodejs-16_x docker-compose ];

  services.openssh = {
      enable = true;
      permitRootLogin = "yes";
  };

  # services.openvpn = {
  #     # You can set openvpn connection
  #     servers = {
  #       privateVPN = {
  #         config = "config /home/nixos/vpn/privatvpn.conf";
  #       };
  #     };
  # };

  virtualisation.docker.enable = true;

  networking.firewall.enable = false;

  # WiFi
  hardware = {
    enableRedistributableFirmware = true;
    firmware = [ pkgs.wireless-regdb ];
  };
  # Networking
  networking = {
    useDHCP = true;
  };

  # put your own configuration here, for example ssh keys:
  users.defaultUserShell = pkgs.fish;
  users.mutableUsers = true;
  users.groups = {
    nixos = {
      gid = 1000;
      name = "nixos";
    };
  };
  users.users = {
    nixos = {
      uid = 1000;
      home = "/home/nixos";
      name = "nixos";
      group = "nixos";
      shell = pkgs.fish;
      extraGroups = [ "wheel" "docker" ];
    };
  };
  users.extraUsers.root.openssh.authorizedKeys.keys = [];
}
