# Example motion 500 divs

<style>
		body {
				margin: 0 !important;
				padding: 0 !important;
		}

		#palco {
			width: 500px;
			height: 300px;
			border: dimgray;
		}

		.animate {
			width: 29px;
			height: 50px;
			position: absolute;
			background-image: url("./small.png");
		}
</style>
<script src="./wasm_exec.js"></script>
<script>
	const go = new Go();
	WebAssembly.instantiateStreaming(fetch("./main.wasm"), go.importObject).then((result) => {
		go.run(result.instance);
	});
</script>
<div id="palco"></div>