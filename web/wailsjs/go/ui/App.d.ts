// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {models} from '../models';
import {player} from '../models';

export function ApiAuthCheck(arg1:string):Promise<{[key: string]: any}>;

export function ApiOnvifDeviceProfile(arg1:string,arg2:string):Promise<{[key: string]: any}>;

export function ApiOnvifDevicePtzMoveAbsolute(arg1:string,arg2:models.PtzControl):Promise<{[key: string]: any}>;

export function ApiOnvifDevicePtzMoveRelative(arg1:string,arg2:models.PtzControl):Promise<{[key: string]: any}>;

export function ApiOnvifDevicePtzStatus(arg1:string,arg2:string):Promise<{[key: string]: any}>;

export function ApiOnvifDeviceStreamurl(arg1:string,arg2:string,arg3:string,arg4:string,arg5:string):Promise<{[key: string]: any}>;

export function ApiOnvifDevices(arg1:string):Promise<{[key: string]: any}>;

export function ApiOnvifPlay(arg1:string,arg2:player.PlayParas):Promise<{[key: string]: any}>;
