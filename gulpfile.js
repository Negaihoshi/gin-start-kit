var gulp = require('gulp');
var inject = require('gulp-inject');
var wiredep = require('wiredep').stream;

gulp.task('default', function() {
  // place code for your default task here
});

gulp.task('bower', function () {
  gulp.src('./templates/layout/_base.html')
    .pipe(wiredep({
      optional: 'configuration',
      goes: 'here'
    }))
    .pipe(gulp.dest('./templates/layout/'));
});

var gulp = require('gulp');
var inject = require('gulp-inject');

gulp.task('test', function () {
  var target = gulp.src('./templates/layout/_base.html');
  // It's not necessary to read the files (will speed up things), we're only after their paths:
  var sources = gulp.src(['./bower_components/**/dist/*.min.js', './bower_components/**/*.min.css'], {read: false});

  return target.pipe(inject(sources))
    .pipe(gulp.dest('./templates/layout/'));
});