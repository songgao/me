window.addEventListener('load', function() {
	var socket = io.connect();
	socket.on('connect', function() {
		var term = new Terminal({
			cols: 80,
			rows: Math.floor((document.getElementById('terminal').clientHeight - 10) / 15),
			screenKeys: true
		});

		term.on('data', function(data) {
			socket.emit('data', data);
		});

		term.open(document.getElementById("terminal"));

		term.write('\x1b[38;5;105m# Got an SSH client? Access this console from your terminal: \x1b[mssh song.gao.io\r\n\n');

		socket.on('data', function(data) {
			term.write(data);
		});

		socket.on('disconnect', function() {
			term.destroy();
		});
	});
}, false);
