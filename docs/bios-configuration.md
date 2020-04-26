# BIOS Configuration

## Steps

<kbd>F10</kbd> (Open Setup)

- _Security_
  - _Create BIOS Administrator Password_ → Set password
- _Advanced_
  - _Boot Options_
    - _USB Storage Boot_ → Disabled
    - _Fast Boot_ → Disabled (disturbs network boot)
    - _Audio Alerts During Boot_ → Disabled (very noisy)
  - _HP Sure Recover_
    - _HP Sure Recover_ → Disabled (tries to restore Windows)
  - _Secure Boot Configuration_
    - _Configure Legacy Support and Secure Boot_ → «Legacy Support Disable and Secure Boot Disable» (**TODO**: use Secure Boot)

<kbd>F10</kbd> → Yes (Save changes & Reboot)

> Enter the code that the BIOS asks to disable the secure boot

## Automation

- Can be automated using Intel vPro/AMT (**TODO**)
- Can be partially automated with a USB programmable keyboard such as :
  - [USB Rubber Ducky](https://shop.hak5.org/products/usb-rubber-ducky-deluxe)
  - [XK-24 USB Programmable Keypad for Windows or Mac](https://www.amazon.com/gp/product/B003MB780E)
  - [Cactus WHID: WiFi HID Injector USB Rubberducky](https://www.tindie.com/products/aprbrother/cactus-whid-wifi-hid-injector-usb-rubberducky/)
