sudo usermod -aG docker $USER &&
newgrp docker &&
sudo chmod +x /usr/local/bin/docker-compose