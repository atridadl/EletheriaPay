<template>
  <b-container>
    <b-alert
      :show="paymentSuccess"
      dismissible
      fade
      variant="success"
    >
      Payment was successful!
    </b-alert>
    <b-alert
      :show="paymentError"
      dismissible
      fade
      variant="danger"
    >
      There was an error making your payment. Please try again.
    </b-alert>
    <img alt="Vue logo" src="/assets/img/logo.png" width="200px">

    <h2>{{ title }}</h2>
    <p>{{ description }}</p>
    <b-row class="h-100 align-items-center">
      <b-col
        xl="5"
        lg="6"
        md="8"
        sm="10"
        class="mx-auto text-center"
      >
        <span>
          <b-form-select
            v-model="currency"
            :options="currencyOptions"
            size="sm"
            class="my-3 radio form-item"
          />
        </span>

        <b-form-group
          id="increments"
        >
          <b-form-radio-group
            id="incrementSelector"
            v-model="amount"
            buttons
            button-variant="theme"
            size="sm"
            name="radio-btn-outline"
            class="form-item"
          >
            <b-form-radio
              class="radio"
              :value="3"
            >
              3
            </b-form-radio>
            <b-form-radio
              class="radio"
              :value="5"
            >
              5
            </b-form-radio>
            <b-form-radio
              class="radio"
              :value="10"
            >
              10
            </b-form-radio>
            <b-form-radio
              class="radio"
              :value="20"
            >
              20
            </b-form-radio>
          </b-form-radio-group>
        </b-form-group>

        <b-form-input
          type="number"
          v-model="amount"
          class="my-2 form-item"
          size="sm"
        />

        <b-form-input
          type="text"
          v-model="donationMessage"
          placeholder="Donation Message"
          class="my-3 form-item"
          size="sm"
        />

        <stripe-element-card
          ref="elementRef"
          v-if="paymentInitiated"
          :pk="publishableKey"
          :elementsOptions="elementsOptions"
          @token="tokenCreated"
        />
        <b-btn
          v-if="paymentInitiated"
          @click="submitPayment"
          class="my-2 form-item"
          size="sm"
          variant="theme"
        >
          Confirm Donation
        </b-btn>
        <b-btn
          v-else
          @click="beginDonation"
          class="my-2 form-item"
          size="sm"
          variant="theme"
        >
          Donate
        </b-btn>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
import { StripeElementCard } from '@vue-stripe/vue-stripe';

export default {
  components: {
    StripeElementCard,
  },
  props: {
    title: String,
    description: String,
    defaultCurrency: String,
  },
  data: () => ({
    amount: 0,
    donationMessage: '',
    publishableKey: '',
    clientSecret: '',
    currency: '',
    elementsOptions: {
      locale: 'auto',
    },
    paymentSuccess: false,
    paymentError: false,
    currencyOptions: [
      { value: 'usd', text: 'USD - United States Dollar' },
      { value: 'cad', text: 'CAD - Canadian Dollar' },
      { value: 'gbp', text: 'GBP - Pounds Sterling' },
      { value: 'aud', text: 'AUD - Australian Dollar' },
      { value: 'eur', text: 'EUR - Euro' },
    ],
    paymentIncrementOptions: [
      { text: '3', value: 3 },
      { text: '5', value: 5 },
      { text: '10', value: 10 },
      { text: '20', value: 20 },
    ],
    paymentType: 0,
    paymentInitiated: false,
  }),
  mounted() {
    this.amount = this.paymentIncrementOptions[1].value;
    console.log(this.defaultCurrency);
    /* eslint-disable */
    const configCurrencyCheck = this.currencyOptions.filter((currency) => currency.value === this.defaultCurrency);
    /* eslint-enable */
    if (configCurrencyCheck.length > 0) {
      this.currency = configCurrencyCheck[0].value;
    } else {
      this.currency = 'usd';
    }
  },
  methods: {
    beginDonation() {
      this.axios.get('/api/pay/config').then((response) => {
        this.publishableKey = response.data.StripePK;
        this.paymentInitiated = true;
      });
    },
    submitPayment() {
      this.$refs.elementRef.submit();
    },
    tokenCreated(token) {
      this.axios.post('/api/pay/secret', {
        amount: this.amount * 100,
        currency: this.currency,
        description: this.donationMessage,
      }).then((response) => {
        this.clientSecret = response.data.client_secret;
      }).then(() => {
        this.$stripe.confirmCardPayment(this.clientSecret, {
          payment_method: {
            card: {
              token: token.id,
            },
          },
        }).then((result) => {
          if (result.error) {
            this.paymentError = 5;
            this.paymentInitiated = false;
            this.clientSecret = '';
            this.publishableKey = '';
            this.donationMessage = '';
          } else if (result.paymentIntent.status === 'succeeded') {
            this.paymentSuccess = 5;
            this.paymentInitiated = false;
            this.clientSecret = '';
            this.publishableKey = '';
            this.donationMessage = '';
          }
        });
      });
    },
  },
};
</script>
