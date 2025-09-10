#!/bin/bash
# build.sh - For Linux builds

set -ex

# Create the target directories
mkdir -p $PREFIX/bin
mkdir -p $PREFIX

# Copy the shared library and module file
cp bin/libplugify-module-golang.so $PREFIX/bin/
cp plugify-module-golang.pmodule $PREFIX/

# Set proper permissions
chmod 755 $PREFIX/bin/libplugify-module-golang.so
chmod 644 $PREFIX/plugify-module-golang.pmodule

# Create activation scripts for proper library path
mkdir -p $PREFIX/etc/conda/activate.d
mkdir -p $PREFIX/etc/conda/deactivate.d

cat > $PREFIX/etc/conda/activate.d/plugify-module-golang.sh << EOF
#!/bin/bash
export PLUGIFY_GO_MODULE_PATH="\${CONDA_PREFIX}:\${PLUGIFY_GO_MODULE_PATH}"
EOF

cat > $PREFIX/etc/conda/deactivate.d/plugify-module-golang.sh << EOF
#!/bin/bash
export PLUGIFY_GO_MODULE_PATH="\${PLUGIFY_GO_MODULE_PATH//\${CONDA_PREFIX}:/}"
EOF

chmod +x $PREFIX/etc/conda/activate.d/plugify-module-golang.sh
chmod +x $PREFIX/etc/conda/deactivate.d/plugify-module-golang.sh