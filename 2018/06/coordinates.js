var path = require('path');
var fs = require('fs');
var readline = require('readline');
var stream = require('stream');

class Point {
  constructor(line) {
    [this.x, this.y] = line.split(',').map(str => parseInt(str));
    this.area = 0;
  }

  distanceFrom(x, y) {
    return Math.abs(x - this.x) + Math.abs(y - this.y);
  }
}

class Grid {
  constructor(points) {
    var size = Math.max(...points.map(point => Math.max(point.x, point.y))) + 1;
    this.grid = Array(size).fill().map(() => Array(size).fill());
    this.points = points;
  }

  fill() {
    this.grid.forEach((row, gy) => {
      row.forEach((_, gx) => {
        var shortestDistance = this.grid.length;
        var closest = [];

        for (var p = 0; p < this.points.length; p++) {
          var point = this.points[p];

          var distance = point.distanceFrom(gx, gy);

          if (distance == shortestDistance) {
            closest.push(p);
          } else if (distance < shortestDistance) {
            shortestDistance = distance;
            closest = [p];

            if (distance == 0) {
              break;
            }
          }
        }

        if (closest.length == 1) {
          row[gx] = closest[0];
        } else {
          row[gx] = -1;
        }
      });
    });

    this.grid.forEach((row, gy) => {
      row.forEach((p, gx) => {
        if (p >= 0 && this.points[p].area != Infinity) {
          if (gy == 0 || gx == 0 || gy == (this.grid.length - 1) || gx == (this.grid.length -1)) {
            this.points[p].area = Infinity;
          } else {
            this.points[p].area++;
          }
        }
      });
    });
  }

  maximumArea() {
    return Math.max(...this.points.filter((point) => point.area != Infinity).map((point) => point.area));
  }
}

function main(argv) {
  var inputPath = argv.length > 1 ? argv[1] : path.join(__dirname, 'input');
  var points = [];
  var instream = fs.createReadStream(inputPath);
  var outstream = new(stream)();
  var rl = readline.createInterface(instream, outstream);

  rl.on('line', function (line) {
    points.push(new Point(line));
  });

  rl.on('close', function () {
    var grid = new Grid(points);

    grid.fill();

    console.log("Maximum area:", grid.maximumArea());
  });
}

main(process.argv.slice(1));
