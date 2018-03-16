function Ghost(id, x, y, img) {
    this.id = id;
    this.row = x / CELL_SIZE;
    this.column = y / CELL_SIZE;

    this.speed = 32;

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
    animation.stop();

    var domElement = this.domElement = this.animation.domElement;

    this.moveRight = function()
    {
        this.dirX = this.speed;
        this.dirY = 0;
        this.animation.rotation = 0;
        if(!this.animation.playing) this.animation.play();
    };

    this.moveLeft = function()
    {
        this.dirX = -this.speed;
        this.dirY = 0;
        this.animation.rotation = 0;
        if(!this.animation.playing) this.animation.play();
    };

    this.moveUp = function()
    {
        this.dirX = 0;
        this.dirY = -this.speed;
        this.animation.rotation = 0;
        if(!this.animation.playing) this.animation.play();
    };

    this.moveDown = function()
    {
        this.dirX = 0;
        this.dirY = this.speed;
        this.animation.rotation = 0;
        if(!this.animation.playing) this.animation.play();
    };

    this.stopMovement = function()
    {
        this.dirX = this.dirY = 0;
        this.animation.stop();
    };

    this.isStopped = function(){
        return this.dirX === 0 && this.dirY === 0;
    };

    this.update = function () {
        this.xp += this.dirX;
        this.yp += this.dirY;

        // console.log("GHOST x : " + this.xp + " : y:" + this.yp);

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