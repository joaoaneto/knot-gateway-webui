var users = require('../models/users');
var DevicesService = require('../services/devices').DevicesService;
var FogService = require('../services/fog').FogService;

var isNotAllowed = function isNotAllowed(device) {
  return !device.allowed;
};

var mapToDevice = function mapToDevice(fogDevice) {
  return {
    uuid: fogDevice.uuid,
    allowed: true,
    name: fogDevice.name,
    online: fogDevice.online
  };
};

var mergeDevices = function mergeDevices(radioDevices, fogDevices) {
  var allowedDevices = fogDevices.map(mapToDevice);
  var notAllowedDevices = radioDevices.filter(isNotAllowed);
  return allowedDevices.concat(notAllowedDevices);
};

var list = function list(req, res, next) {
  var devicesSvc = new DevicesService();
  var fogSvc = new FogService();
  users.getUserByUUID(req.user.uuid, function onUser(userErr, user) {
    if (userErr) {
      next(userErr);
    } else {
      devicesSvc.list(function onDevicesReturned(deviceErr, radioDevices) {
        if (deviceErr) {
          next(deviceErr);
        } else {
          fogSvc.getDevices(user, function onFogDevicesReturned(fogErr, fogDevices) {
            var devices;
            if (fogErr) {
              next(fogErr);
            } else {
              devices = mergeDevices(radioDevices, fogDevices);
              res.json(devices);
            }
          });
        }
      });
    }
  });
};

var add = function add(req, res, next) {
  var devicesSvc = new DevicesService();
  var device = {
    mac: req.params.id,
    name: req.body.name,
    allowed: req.body.allowed
  };
  devicesSvc.update(device, function onDevicesUpdated(devicesErr, updated) {
    if (devicesErr) {
      next(devicesErr);
    } else if (!updated) {
      res.sendStatus(500); // TODO: verify in which case a device isn't updated
    } else {
      res.end();
    }
  });
};

var remove = function remove(req, res, next) {
  var fogSvc = new FogService();
  users.getUserByUUID(req.user.uuid, function onUser(userErr, user) {
    if (userErr) {
      next(userErr);
    } else {
      fogSvc.removeDevice(user, req.params.id, function onDeviceRemoved(fogErr) {
        if (fogErr) {
          next(fogErr);
        } else {
          res.end();
        }
      });
    }
  });
};

var update = function update(req, res, next) {
  if (req.body.allowed) {
    add(req, res, next);
  } else {
    remove(req, res, next);
  }
};

module.exports = {
  list: list,
  update: update
};
