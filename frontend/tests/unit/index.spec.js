import { shallowMount } from '@vue/test-utils';
import Index from '@/components/Index.vue';
import BootstrapVue from 'bootstrap-vue';
import { createLocalVue } from '@vue/test-utils'
import Vue from 'vue';

// create an extended `Vue` constructor
const localVue = createLocalVue()

// install plugins as normal
localVue.use(BootstrapVue)

describe('Index.vue', () => {
  it('Index.vue - title and description render', async () => {
    const title = 'Title';
    const description = 'Description';
    const wrapper = shallowMount(Index, {
      propsData: { title, description },
      localVue
    });
    expect(wrapper.text()).toMatch(title);
    expect(wrapper.text()).toMatch(description);
    expect(wrapper.find('#nonAnnualIncrements').exists()).toBeTruthy();
    expect(wrapper.find('#nonAnnualIncrements').exists()).toBeTruthy();
    await wrapper.setData({paymentType: 2});
    expect(wrapper.vm.paymentType).toBe(2);
    expect(wrapper.find('#annualIncrements').exists()).toBeTruthy();
  });

  it('Index.vue - payment increments work correctly', async () => {
    const enableOneTime = true;
    const enableMonthly = true;
    const enableAnnual = true;
    const wrapper = shallowMount(Index, {
      propsData: { enableOneTime, enableMonthly, enableAnnual },
      localVue
    });
    expect(wrapper.find('#nonAnnualIncrements').exists()).toBeTruthy();
    expect(wrapper.find('#nonAnnualIncrements').exists()).toBeTruthy();
    await wrapper.setData({paymentType: 2});
    expect(wrapper.vm.paymentType).toBe(2);
    expect(wrapper.find('#annualIncrements').exists()).toBeTruthy();
  });

  it('Index.vue - monthly payment type disabled', async () => {
    const enableOneTime = true;
    const enableMonthly = false;
    const enableAnnual = true;
    const wrapper = shallowMount(Index, {
      propsData: { enableOneTime, enableMonthly, enableAnnual },
      localVue
    });
    expect(wrapper.find('#oneTimePaymentSelector').exists()).toBeTruthy();
    expect(!wrapper.find('#monthlyPaymentSelector').exists()).toBeTruthy();
    expect(wrapper.find('#annualPaymentSelector').exists()).toBeTruthy();
  });

  it('Index.vue - annual payment type disabled', async () => {
    const enableOneTime = true;
    const enableMonthly = true;
    const enableAnnual = false;
    const wrapper = shallowMount(Index, {
      propsData: { enableOneTime, enableMonthly, enableAnnual },
      localVue
    });
    expect(wrapper.find('#oneTimePaymentSelector').exists()).toBeTruthy();
    expect(wrapper.find('#monthlyPaymentSelector').exists()).toBeTruthy();
    expect(!wrapper.find('#annualPaymentSelector').exists()).toBeTruthy();
  });

  it('Index.vue - one-time payment type only', async () => {
    const enableOneTime = true;
    const enableMonthly = false;
    const enableAnnual = false;
    const wrapper = shallowMount(Index, {
      propsData: { enableOneTime, enableMonthly, enableAnnual },
      localVue
    });
    expect(wrapper.find('#oneTimePaymentSelector').exists()).toBeTruthy();
    expect(!wrapper.find('#monthlyPaymentSelector').exists()).toBeTruthy();
    expect(!wrapper.find('#annualPaymentSelector').exists()).toBeTruthy();
  });

  it('Index.vue - all payment tupes disabled', async () => {
    const enableOneTime = false;
    const enableMonthly = false;
    const enableAnnual = false;
    const wrapper = shallowMount(Index, {
      propsData: { enableOneTime, enableMonthly, enableAnnual },
      localVue
    });
    expect(!wrapper.find('#oneTimePaymentSelector').exists()).toBeTruthy();
    expect(!wrapper.find('#monthlyPaymentSelector').exists()).toBeTruthy();
    expect(!wrapper.find('#annualPaymentSelector').exists()).toBeTruthy();
  });
});
