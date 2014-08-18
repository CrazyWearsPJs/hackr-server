var gulp = require('gulp');
var gutil = require('gulp-util');

var sass = require('gulp-sass');

var browserSync = require('browser-sync');

var uglify = require('gulp-uglify');
var concat = require('gulp-concat');
var rename = require('gulp-rename');
var clean = require('gulp-clean');

var jshint = require('gulp-jshint')
var stylish = require('jshint-stylish');

var ngAnnotate = require('gulp-ng-annotate');

var runSequence = require('run-sequence');

var sources = {
    js: ['./app/components/**/*.js','./app/app.js'],
    sass: './app/styles/**/*.scss',
    html: './app/**/*.html'
};

var destinations = {
    css : './app/css',
    js : './app/js'
};

gulp.task('brower-sync', function(){
    browserSync.init(null, {
        open: false,
        server: {
            baseDir: "./app"
        },
        watchOptions: {
            debounceDelay: 1000
        }
    });
});

gulp.task('style', function(){
    gulp.src(sources.sass)
        .pipe(sass({outputStyle: 'compressed', errLogToConsole: true}))
        .pipe(gulp.dest(destinations.css));
});

gulp.task('src', function(){
    gulp.src(sources.js, {base: 'app/'})
        .pipe(concat('all.js'))
        .pipe(ngAnnotate())
        .pipe(gulp.dest(destinations.js))
        .pipe(uglify())
        .pipe(rename('all.min.js'))
        .pipe(gulp.dest(destinations.js));
});

gulp.task('lint', function(){
    gulp.src(sources.js,{base: 'app/'})
        .pipe(jshint())
        .pipe(jshint.reporter(stylish));
});

gulp.task('watch', function(){
    gulp.watch(sources.sass, ['style']);
    gulp.watch(sources.js, ['lint', 'src']);
});

gulp.watch('./app/**/**', function(file){
    if (file.type === "changed"){
        browserSync.reload(file.path);
    }
});

gulp.task('clean', function(){
    gulp.src([destinations.js, destinations.css], {read:false})
        .pipe(clean());
});

gulp.task('build', function(){
    runSequence('clean', 'lint', ['style', 'src']);
});

gulp.task('default', ['build', 'brower-sync', 'watch']);
