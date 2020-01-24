#!/bin/sh

if [ $APP_ENV == "local" ]; then
    cp /go/src/api/.env.local /go/src/api/.env;
    sed -i "s/\$SERVER_NAME/beego.local/g" /etc/nginx/sites-enabled/app.conf;
    sed -i "s/\$REDIRECT_URL/beego.local/g" /etc/nginx/sites-enabled/app.conf;
    sed -i "s/\$HTTP_X_FORWARDED_HOST/beego.local/g" /etc/nginx/sites-enabled/app.conf;
fi

if [ $APP_ENV == "alpha" ]; then
    mv /app/.env.alpha /app/.env;
    sed -i "s/\$SERVER_NAME/alpha-beego.com/g" /etc/nginx/sites-enabled/app.conf;
    sed -i "s/\$REDIRECT_URL/alpha-beego.com/g" /etc/nginx/sites-enabled/app.conf;
    sed -i "s/\$HTTP_X_FORWARDED_HOST/alpha-beego.com/g" /etc/nginx/sites-enabled/app.conf;
fi

if [ $APP_ENV == "staging" ]; then
    mv /app/.env.staging /app/.env;
    sed -i "s/\$SERVER_NAME/staging-beego.com/g" /etc/nginx/sites-enabled/app.conf;
    sed -i "s/\$REDIRECT_URL/staging-beego.com/g" /etc/nginx/sites-enabled/app.conf;
    sed -i "s/\$HTTP_X_FORWARDED_HOST/staging-beego.com/g" /etc/nginx/sites-enabled/app.conf;
fi

if [ $APP_ENV == "preprod" ]; then
    sed -i "s/\$SERVER_NAME/preprod-beego.com/g" /etc/nginx/sites-enabled/app.conf;
    sed -i "s/\$REDIRECT_URL/preprod-beego.com/g" /etc/nginx/sites-enabled/app.conf;
    sed -i "s/\$HTTP_X_FORWARDED_HOST/preprod-beego.com/g" /etc/nginx/sites-enabled/app.conf;
fi

if [ $APP_ENV == "production" ]; then
    sed -i "s/\$SERVER_NAME/beego.com/g" /etc/nginx/sites-enabled/app.conf;
    sed -i "s/\$REDIRECT_URL/beego.com/g" /etc/nginx/sites-enabled/app.conf;
    sed -i "s/\$HTTP_X_FORWARDED_HOST/beego.com/g" /etc/nginx/sites-enabled/app.conf;
fi

supervisord -n -c /etc/supervisord.conf