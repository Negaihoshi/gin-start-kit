var gulp = require('gulp');
var wiredep = require('wiredep').stream;

gulp.task('bower', function() {
  gulp.src('./templates/layout.tmpl')
    .pipe(wiredep({
      optional: 'configuration',
      goes: 'here',
    }))
    .pipe(gulp.dest('./templates/'));
});

gulp.task('default', ['bower']);
