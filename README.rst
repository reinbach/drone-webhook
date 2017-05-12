# Drone Webhook Plugin

Basic webhook plugin docker container that works with [Drone](https://github.com/drone/drone)
A copy from [custom go plugin example](http://docs.drone.io/creating-custom-plugins-golang/) with minor tweaks.


## Usage

    pipeline:
      webhook:
        image: reinbach/drone-webhook
        url: http://foo.com
        method: post
        body: hello world
