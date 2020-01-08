import SocketTask from './socketTask';
import { getParameterError, shouleBeObject } from './utils';

let socketTasks = [];
let socketsCounter = 1;

function connectSocket(options) {
    const name = 'connectSocket';

    return new Promise((resolve, reject) => {
        // options must be an Object
        const isObject = shouleBeObject(options);
        if(!isObject.res) {
            const res = { errMsg: `${name}${isObject.msg}` };
            console.error(res.errMsg);
            return reject(res);
        }

        const {
            url,
            protocols,
            success,
            fail,
            complete,
        } = options;
        const res = { errMsg: 'connectSocket:ok' };

        // options.url must be String
        if(typeof url !== 'string') {
            res.errMsg = getParameterError({
                name,
                para: 'url',
                correct: 'String',
                wrong: url
            });
            console.error(res.errMsg);
            typeof fail === 'function' && fail(res);
            typeof complete === 'function' && complete(res);
            return reject(res);
        }

        // options.url must be invalid
        if(!url.startsWith('ws://') && !url.startsWith('wss://')) {
            res.errMsg = `request:fail invalid url "${url}"`;
            console.error(res.errMsg);
            typeof fail === 'function' && fail(res);
            typeof complete === 'function' && complete(res);
            return reject(res);
        }

        // protocols must be array
        const protocolsArr = Array.isArray(protocols) ? protocols : null;

        // 2 connection at most
        // if(socketTasks.length > 1) {
        //     res.errMsg = '同时最多发起 2 个 socket 请求，更多请参考文档。';
        //     console.error(res.errMsg);
        //     typeof fail === 'function' && fail(res);
        //     typeof complete === 'function' && complete(res);
        //     return reject(res);
        // }

        const task = new SocketTask(url, protocolsArr, options.binaryType);
        task.destroyWhenClose = function () { // eslint-disable-line
            socketTasks = socketTasks.filter(socketTask => socketTask !== this);
        };
        socketTasks.push(task);

        res.socketTaskId = socketsCounter++;
        typeof success === 'function' && success(res);
        typeof complete === 'function' && complete(res);

        return resolve(task);
    });
}

export default {
    connectSocket,
};
