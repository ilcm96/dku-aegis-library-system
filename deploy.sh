#!/bin/bash
set -e

echo "CURRENT SERVER: Oracel Arm Seoul"

cd /home/ubuntu/docker/dku-aegis-library-system

CURRENT_COLOR=$(curl -s https://dku-aegis-library-system.seongmin.dev/color)

if [ "$CURRENT_COLOR" == "blue" ]; then
  NEW_COLOR="green"
else
  NEW_COLOR="blue"
fi

echo "Current color is $CURRENT_COLOR. Deploying to $NEW_COLOR..."

docker compose -f compose-app.yml pull
docker compose -f compose-app.yml up -d dku-${NEW_COLOR}-app-1 dku-${NEW_COLOR}-app-2

echo "New version deployed to $NEW_COLOR. Checking health..."

# Here you should add health check logic for the new version
# This example assumes the health check endpoint is /health
HEALTHY=0
for i in {1..10}; do
  if docker compose -f compose-app.yml exec -T dku-nginx bash -c "curl -s http://dku-${NEW_COLOR}-app-1:3000/health" | grep -q "OK"; then
    HEALTHY=$((HEALTHY + 1))
  fi
  if docker compose -f compose-app.yml exec -T dku-nginx bash -c "curl -s http://dku-${NEW_COLOR}-app-2:3000/health" | grep -q "OK"; then
    HEALTHY=$((HEALTHY + 1))
  fi
  if [ "$HEALTHY" -ge 2 ]; then
    break
  fi
  echo "Waiting for the new version to become healthy..."
  sleep 3
done

if [ "$HEALTHY" -lt 2 ]; then
  echo "New version is not healthy. Aborting."
  docker compose -f compose-app.yml stop ${NEW_COLOR}-app-1 ${NEW_COLOR}-app-2
  exit 1
fi

echo "New version is healthy. Switching traffic to $NEW_COLOR..."

# Update Nginx environment variable and reload
# Not work on alpine image due to different behavior of cp -f
docker compose -f compose-app.yml exec -T dku-nginx bash -c "sed "s/${CURRENT_COLOR}/${NEW_COLOR}/g" /etc/nginx/nginx.conf > /etc/nginx/nginx.conf.tmp \
&& cp -f /etc/nginx/nginx.conf.tmp /etc/nginx/nginx.conf \
&& rm /etc/nginx/nginx.conf.tmp \
&& nginx -s reload"

echo "Traffic switched to $NEW_COLOR. Stopping old version..."

# Stop the old version
docker compose -f compose-app.yml stop dku-${CURRENT_COLOR}-app-1 dku-${CURRENT_COLOR}-app-2

echo "Deployment complete. Now serving $NEW_COLOR."
