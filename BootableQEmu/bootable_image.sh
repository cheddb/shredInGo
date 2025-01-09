#!/bin/bash
set -e
# Ensure the script runs with superuser privileges
if [ "$EUID" -ne 0 ]; then
  echo "Please run as root"
  exit 1
fi

# Variables
KERNEL_NAME="linux-6.11"
KERNEL_ARCHIVE="$KERNEL_NAME.tar.xz"
KERNEL_URL="https://cdn.kernel.org/pub/linux/kernel/v6.x/$KERNEL_ARCHIVE"
IMAGE_FILE="rootfs.img"
IMAGE_SIZE=256M
MOUNT_POINT="mnt"

# Install necessary tools
apt-get update
apt-get install -y qemu-system-x86 build-essential libncurses-dev bison flex libssl-dev libelf-dev

# Download and extract the kernel
curl -o $KERNEL_ARCHIVE $KERNEL_URL

if [ ! -f "$KERNEL_ARCHIVE" ]; then
  echo "Kernel archieve missing"
  exit 1
fi

tar -xvf $KERNEL_ARCHIVE
cd $KERNEL_NAME

# Configure and build the kernel
make defconfig
make -j$(nproc)
cd ..

# Create filesystem image
fallocate -l $IMAGE_SIZE $IMAGE_FILE
mkfs.ext4 $IMAGE_FILE
mkdir -p $MOUNT_POINT
sudo mount -o loop $IMAGE_FILE $MOUNT_POINT

# Setup filesystem
sudo debootstrap --variant=minbase --arch=amd64 focal $MOUNT_POINT
cat <<EOL > mnt/init
#!/bin/sh
echo "hello world"
exec /bin/sh
EOL

sudo chmod +x mnt/init
sudo umount $MOUNT_POINT

# Launch QEMU
qemu-system-x86_64 -kernel $KERNEL_NAME/arch/x86/boot/bzImage \
  -append "console=ttyS0 root=/dev/sda rw init=/init" \
  -hda $IMAGE_FILE -nographic

