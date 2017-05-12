Drone Webhook Plugin
====================

Basic webhook plugin docker container that works with `Drone <https://github.com/drone/drone>`_.

A copy from `custom go plugin example <http://docs.drone.io/creating-custom-plugins-golang/>`_ with minor tweaks.


Usage
-----

.. code-block:: yaml

    pipeline:
      webhook:
        image: reinbach/drone-webhook
        url: http://foo.com
        method: post
        body: hello world
