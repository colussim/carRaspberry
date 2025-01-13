#!/bin/bash

# Path to service files (change this if necessary)
SERVICE_DIR=/etc/systemd/system
AMPLI_POWER_SERVICE=services/amplipower.service
AMPLI_STOP_SERVICE=services/amplipower-stop.service

echo "Copy amplipower binary ..."
cp -p amplipower /usr/local/bin/amplipower

echo "Installing services..."

# Copy service files to /etc/systemd/system
cp $AMPLI_POWER_SERVICE $SERVICE_DIR/
cp $AMPLI_STOP_SERVICE $SERVICE_DIR/

# Reload systemd configuration files 
sudo systemctl daemon-reload

# Enable services at startup  
sudo systemctl enable amplipower.service
sudo systemctl enable amplipower-stop.service


