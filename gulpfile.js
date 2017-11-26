'use strict'

var gulp = require('gulp');
var sass = require('gulp-sass');
var uglify = require('gulp-uglify');
var cleanCSS = require('gulp-clean-css');
var htmlmin = require('gulp-htmlmin');

var paths = {
    src: {
        style: './scss/**/*.scss',
        scripts: './scripts/*.js',
        html: './templates/**/*.html'
    },
    dest: {
        style: './public/style',
        scripts: './public/js',
        html: './templates_minified',
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
});

gulp.task('js', function() {
	return gulp.src(paths.src.scripts)
        .pipe(uglify())
        .on('error', errorHandler)
        .pipe(gulp.dest(paths.dest.scripts));
});

gulp.task('htmlmin', function () {
    return gulp.src(paths.src.html)
        .pipe(htmlmin({collapseWhitespace: true}))
        .pipe(gulp.dest(paths.dest.html));
});


gulp.task('watch', function () {
    gulp.watch(paths.src.style, ['css'])
    gulp.watch(paths.src.scripts, ['js'])
    gulp.watch(paths.src.html, ['htmlmin'])
});

gulp.task('default', ['css', 'js', 'htmlmin']);