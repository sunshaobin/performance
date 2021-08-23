var r = 0;
const t1 = new Date().getTime();

for (var c = 0; c < 1000000000; c++) {
    r += c;
}
const t2 = new Date().getTime();

console.log("t1",t1);
console.log("t2",t2);
console.log(t2-t1);