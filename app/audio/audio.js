let createNoise = require("audio-noise");
const write = require("web-audio-write")();

// let noise = createNoise("pink");

//create array filled with pink noise
// let arr = noise(Array(1024));

// const noise = require("./")({ color: "white", format: "stereo audiobuffer" });
const noise = createNoise({ color: "white", format: "stereo audiobuffer" });

(function tick(err) {
  if (err) throw err;
  write(noise(), tick);
})();

// var Generator = require("audio-generator/stream");
// var Speaker = require("audio-speaker/stream");

// Generator(
//   //Generator function, returns sample values -1..1 for channels
//   function(time) {
//     return [
//       Math.sin(Math.PI * 2 * time * 439), //channel 1
//       Math.sin(Math.PI * 2 * time * 441) //channel 2
//     ];
//   },
// )

// //   {
// //     //Duration of generated stream, in seconds, after which stream will end.
// //     duration: 5, //Infinity,

// //     //Periodicity of the time.
// //     period: 2 //Infinity
// //   }
// // )
// //   .on("error", function(e) {
// //     //error happened during generation the frame
// //   })
// //   .pipe(Speaker());

// // function alertTerminal() {
// //   console.log("\007");
// // }

// // // var beep = require("node-beep");
// // // beep(1);

// // var done = (function wait() {
// //   if (!done) setTimeout(wait, 1000);
// // })();

// // function sleep(ms) {
// //   return new Promise(resolve => {
// //     setTimeout(resolve, ms);
// //   });
// // }
// // async function init() {
// //   await sleep(3000);
// //   process.stdout.write("\x07");
// //   // beep(1);
// //   console.log("exiting");
// //   done = true;
// // }

// // init();
