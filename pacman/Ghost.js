function Ghost(x, y, img) {
    this.row = x / CELL_SIZE;
    this.column = y / CELL_SIZE;

    this.speed = 1;

    var xp = this.xp = this.startX = x;
    var yp = this.yp = this.startY = y;

    var dirX = this.dirX = 0;
    var dirY = this.dirY = 0;

    this.width = 32;
    this.height = 32;

    var animation = this.animation = new Animation(img, 32, 32, domElement);
    animation.offsetY = 0;
    animation.offsetX = 0;
    animation.looping = true;
    animation.play();

    var domElement = this.domElement = this.animation.domElement;

    this.update = function () {
        this.xp += this.dirX;
        this.yp += this.dirY;

        console.log("GHOST x : " + this.xp + " : y:" + this.yp);

        this.animation.x = this.xp;
        this.animation.y = this.yp;
        this.animation.update();
    };

    this.render = function () {
        this.animation.render();
    };

    this.reset = function () {
        this.xp = this.startX;
        this.yp = this.startY;
    }

    this.init = function () {
        this.update();
        this.render();
    }
}