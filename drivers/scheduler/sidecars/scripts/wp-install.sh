#!/usr/bin/env bash

function wait_for_mysql() {
    until nc -z -v -w30 ${WORDPRESS_DB_HOST} 3306
    do
      echo ":: Waiting for database connection..."
      sleep 5
    done
}

function wordpress_install() {
    wp core install --url="${WORDPRESS_URL}" --title="TestWordPress" --admin_user="admin" --admin_password="correcthorsebatterystaple" --admin_email="noreply@portworx.com" --skip-email
}

cd /wordpress
wait_for_mysql

echo ":: Downloading wordpress..."
wp core download

echo ":: Generating wordpress config..."
wp config create --dbname=pwx --dbhost="${WORDPRESS_DB_HOST}" --dbuser="root" --dbpass="${WORDPRESS_DB_PASSWORD}" --force

echo ":: Generating wordpress database..."
wp db create --dbuser=root --dbpass="${WORDPRESS_DB_PASSWORD}"

echo ":: Installing wordpress..."
wordpress_install

echo ":: Generating junk posts..."
wp post generate --count=100

echo ":: Installing plugins..."
for plugin_name in ${WORDPRESS_PLUGINS//,/ }
do
    wp plugin install "${plugin_name}" --activate
done

echo ":: Installing theme..."
wp theme install "${WORDPRESS_THEME}" --activate

echo ":: Allowing wordpress pods to come online"
touch /wordpress/installed

echo ":: Sleeping forever. Exec into me for debugging"
while true; do sleep 30; done;
