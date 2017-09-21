'use strict'

var gulp = require('gulp')
var sass = require('gulp-sass')
var paths = {
    src: {
        style: './scss/**/*.scss'
    },
    dest: {
        style: './public/style'
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
        .on('error', errorHandler)
        .pipe(gulp.dest(paths.dest.style))
})

gulp.task('watch', function () {
    gulp.watch(paths.src.style, )
})

gulp.task('default', ['css'])