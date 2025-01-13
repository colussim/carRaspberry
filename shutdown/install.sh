#!/bin/bash

# Path to service files (change this if necessary)
SERVICE_DIR=/etc/systemd/system
BUTTON_POWER_SERVICE=services/button_shutdown.service

echo "Copy button_shutdown binary ..."
cp -p button_shutdown /usr/local/bin/button_shutdown

echo "Installing services..."

# Copy service files to /etc/systemd/system
cp $BUTTON_POWER_SERVICE $SERVICE_DIR/

# Reload systemd configuration files 
sudo systemctl daemon-reload

# Enable services at startup  
sudo systemctl enable button_shutdown.service

