"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildServices = void 0;
var _1 = require("./");
var buildServices = function (provider) {
    return {
        pet: new _1.PetService(provider),
        store: new _1.StoreService(provider),
        user: new _1.UserService(provider),
    };
};
exports.buildServices = buildServices;
