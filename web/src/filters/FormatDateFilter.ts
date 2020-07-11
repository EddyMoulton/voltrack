import Vue from 'vue';
import moment from 'moment';

export default Vue.filter('formatDate', function (value: Date, format = '') {
  if (value) {
    if (!format) {
      format = 'DD/MM/YYYY hh:mm';
    }

    return moment(String(value)).format(format);
  }
});
