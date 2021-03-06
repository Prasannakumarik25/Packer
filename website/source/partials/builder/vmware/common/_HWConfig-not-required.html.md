<!-- Code generated from the comments of the HWConfig struct in builder/vmware/common/hw_config.go; DO NOT EDIT MANUALLY -->

-   `cpus` (int) - The number of cpus to use when building the VM.
    
-   `memory` (int) - The amount of memory to use when building the VM in megabytes.
    
-   `cores` (int) - The number of cores per socket to use when building the VM. This
    corresponds to the cpuid.coresPerSocket option in the .vmx file.
    
-   `network` (string) - This is the network type that the virtual machine will be created with.
    This can be one of the generic values that map to a device such as
    hostonly, nat, or bridged. If the network is not one of these values,
    then it is assumed to be a VMware network device. (VMnet0..x)
    
-   `network_adapter_type` (string) - This is the ethernet adapter type the the virtual machine will be
    created with. By default the `e1000` network adapter type will be used
    by Packer. For more information, please consult [Choosing a network
    adapter for your virtual
    machine](https://kb.vmware.com/s/article/1001805) for desktop VMware
    clients. For ESXi, refer to the proper ESXi documentation.
    
-   `sound` (bool) - Specify whether to enable VMware's virtual soundcard device when
    building the VM. Defaults to false.
    
-   `usb` (bool) - Enable VMware's USB bus when building the guest VM. Defaults to false.
    To enable usage of the XHCI bus for USB 3 (5 Gbit/s), one can use the
    vmx_data option to enable it by specifying true for the usb_xhci.present
    property.
    
-   `serial` (string) - This specifies a serial port to add to the VM. It has a format of
    `Type:option1,option2,...`. The field `Type` can be one of the following
    values: `FILE`, `DEVICE`, `PIPE`, `AUTO`, or `NONE`.
    
    * `FILE:path(,yield)` - Specifies the path to the local file to be used
      as the serial port.
    
        * `yield` (bool) - This is an optional boolean that specifies
          whether the vm should yield the cpu when polling the port. By
          default, the builder will assume this as `FALSE`.
    
    * `DEVICE:path(,yield)` - Specifies the path to the local device to be
    ??  used as the serial port. If `path` is empty, then default to the first
      serial port.
    
        * `yield` (bool) - This is an optional boolean that specifies
          whether the vm should yield the cpu when polling the port. By
          default, the builder will assume this as `FALSE`.
    
    * `PIPE:path,endpoint,host(,yield)` - Specifies to use the named-pipe
      "path" as a serial port. This has a few options that determine how the
      VM should use the named-pipe.
    
        * `endpoint` (string) - Chooses the type of the VM-end, which can be
          either a `client` or `server`.
    
        * `host` (string)     - Chooses the type of the host-end, which can
          be either `app` (application) or `vm` (another virtual-machine).
    
        * `yield` (bool)      - This is an optional boolean that specifies
          whether the vm should yield the cpu when polling the port. By
          default, the builder will assume this as `FALSE`.
    
    * `AUTO:(yield)` - Specifies to use auto-detection to determine the
      serial port to use. This has one option to determine how the VM should
      support the serial port.
    
        * `yield` (bool) - This is an optional boolean that specifies
          whether the vm should yield the cpu when polling the port. By
          default, the builder will assume this as `FALSE`.
    
    * `NONE` - Specifies to not use a serial port. (default)
    
-   `parallel` (string) - This specifies a parallel port to add to the VM. It has the format of
    `Type:option1,option2,...`. Type can be one of the following values:
    `FILE`, `DEVICE`, `AUTO`, or `NONE`.
    
    * `FILE:path` 		- Specifies the path to the local file to be used
      for the parallel port.
    
    * `DEVICE:path`	 	- Specifies the path to the local device to be used
      for the parallel port.
    
    * `AUTO:direction`   - Specifies to use auto-detection to determine the
      parallel port. Direction can be `BI` to specify bidirectional
      communication or `UNI` to specify unidirectional communication.
    
    * `NONE` 			- Specifies to not use a parallel port. (default)
    