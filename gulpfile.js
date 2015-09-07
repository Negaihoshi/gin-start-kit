var gulp = require('gulp');
var inject = require('gulp-inject');
var wiredep = require('wiredep').stream;
var jade = require('gulp-jade');

gulp.task('default', function() {
  // place code for your default task here
});

gulp.task('bower', function () {
  gulp.src('./templates/layout.jade')
    .pipe(wiredep({
      optional: 'configuration',
      goes: 'here'
    }))
    .pipe(gulp.dest('./templates/'));
});



gulp.task('templates', function() {

  gulp.src('./templates/*.jade')
    .pipe(jade({
      pretty: true
    }))
    .pipe(gulp.dest('./.tmp/public/template/'));
});

gulp.task('default', ['bower', 'templates']);
