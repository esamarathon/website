'use strict'

var gulp = require('gulp')
var sass = require('gulp-sass')
var uglify = require('gulp-uglify')
let cleanCSS = require('gulp-clean-css')

var paths = {
    src: {
        style: './scss/**/*.scss',
        scripts: './scripts/*.js'
    },
    dest: {
        style: './public/style',
        scripts: './public/js'
    }
}

function errorHandler (error) {
    console.log('\n\n\n----------------------------------------------------\n Begin error\n');
    console.error(error.toString());
    console.log('\n\n\n----------------------------------------------------\n end of error\n');
    this.emit('end');
}

gulp.task('css', function () {
    return gulp.src(paths.src.style)
        .pipe(sass().on('error', sass.logError))
        .pipe(cleanCSS({compatibility: 'ie9'}))
        .on('error', errorHandler)
        .pipe(gulp.dest(paths.dest.style))
})

gulp.task('js', function() {
	gulp.src(paths.src.scripts)
	.pipe(uglify())
    .on('error', errorHandler)
	.pipe(gulp.dest(paths.dest.scripts));
});


gulp.task('watch', function () {
    gulp.watch(paths.src.style, ['css', 'js'])
})

gulp.task('default', ['css', 'js'])