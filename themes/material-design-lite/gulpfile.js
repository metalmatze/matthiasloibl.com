var gulp = require('gulp');
var minifyCss = require('gulp-minify-css');
var sass = require('gulp-sass');
var uglify = require('gulp-uglify');
var concat = require('gulp-concat');

gulp.task('default', [
    'js',
    'css',
    'concat',
    'copy'
]);

gulp.task('build', [
    'minify'
]);

gulp.task('watch', ['default'], function () {
    gulp.watch('assets/**/*.js', ['js']);
    gulp.watch('assets/**/*.scss', ['css']);
});

gulp.task('js', function () {
});

gulp.task('css', function () {
    gulp.src('assets/scss/*.scss')
        .pipe(sass().on('error', sass.logError))
        .pipe(gulp.dest('./static/css'));
});

gulp.task('concat', function () {
    return gulp.src([
            'node_modules/particles.js/particles.js',
            'assets/js/particles.js'
        ])
        .pipe(concat('particles.js'))
        .pipe(gulp.dest('static/js'));
});

gulp.task('minify', ['default'], function () {
    //gulp.src('static/app.js')
    //    .pipe(uglify())
    //    .pipe(gulp.dest('static'));
    gulp.src('static/css/main.css')
        .pipe(minifyCss())
        .pipe(gulp.dest('static/css'));
});

gulp.task('copy', function () {
    gulp.src('node_modules/material-design-lite/material.min.js')
        .pipe(gulp.dest('static/js'));
    gulp.src('node_modules/mobile-detect/mobile-detect.min.js')
        .pipe(gulp.dest('static/js'));
});