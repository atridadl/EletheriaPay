import '@babel/polyfill';
import 'mutationobserver-shim';
import '@stripe/stripe-js';
import Vue from 'vue';
import { StripePlugin } from '@vue-stripe/vue-stripe';
import axios from 'axios';
import VueAxios from 'vue-axios';
import './plugins/bootstrap-vue';
import App from './App.vue';

Vue.config.productionTip = false;

axios.get('/api/pay/config').then((response) => {
  const options = {
    pk: response.data.StripePK,
    stripeAccount: response.data.StripeACC,
    apiVersion: response.data.StripeApiVersion,
    locale: response.data.StripeLocale,
  };
  Vue.use(StripePlugin, options);
});

Vue.use(VueAxios, axios);

new Vue({
  render: (h) => h(App),
}).$mount('#app');
