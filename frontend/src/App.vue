<template>
  <div
    id="app"
  >
    <b-navbar
      id="top-nav"
    >
      <b-navbar-nav>
        <b-nav-item
          :href="backLink"
        >
          <arrow-left-circle-icon
            size="1.5x"
            class="navIcons"
          />
        </b-nav-item>
      </b-navbar-nav>
      <b-navbar-nav class="ml-auto">
        <b-nav-item
          v-if="emailLink !== ''"
          :href="'mailto:' + emailLink"
        >
          <mail-icon
            size="1.5x"
            class="navIcons"
          />
        </b-nav-item>
        <b-nav-item
          v-if="gitLink !== ''"
          :href="gitLink"
        >
          <git-pull-request-icon
            size="1.5x"
            class="navIcons"
          />
        </b-nav-item>
        <b-nav-item
          v-if="githubLink !== ''"
          :href="githubLink"
        >
          <github-icon
            size="1.5x"
            class="navIcons"
          />
        </b-nav-item>
        <b-nav-item
          v-if="gitlabLink !== ''"
          :href="gitlabLink"
        >
          <gitlab-icon
            size="1.5x"
            class="navIcons"
          />
        </b-nav-item>
        <b-nav-item
          v-if="facebookLink !== ''"
          :href="facebookLink"
        >
          <facebook-icon
            size="1.5x"
            class="navIcons"
          />
        </b-nav-item>
        <b-nav-item
          v-if="twitterLink !== ''"
          :href="twitterLink"
        >
          <twitter-icon
            size="1.5x"
            class="navIcons"
          />
        </b-nav-item>
        <b-nav-item
          v-if="instagramLink !== ''"
          :href="instagramLink"
        >
          <instagram-icon
            size="1.5x"
            class="navIcons"
          />
        </b-nav-item>
        <b-nav-item
          v-if="linkedinLink !== ''"
          :href="linkedinLink"
        >
          <linkedin-icon
            size="1.5x"
            class="navIcons"
          />
        </b-nav-item>
        <b-nav-item
          v-if="theme == 'light'"
        >
          <moon-icon
            @click="toggleTheme()"
            size="1.5x"
            class="navIcons"
          />
        </b-nav-item>
        <b-nav-item
          v-if="theme == 'dark'"
        >
          <sun-icon
            @click="toggleTheme()"
            size="1.5x"
            class="navIcons"
          />
        </b-nav-item>
      </b-navbar-nav>
    </b-navbar>
    <Index
      v-if="configLoaded"
      :title="title"
      :description="description"
      :defaultCurrency="defaultCurrency"
    />
  </div>
</template>
<script>
import {
  SunIcon,
  MoonIcon,
  ArrowLeftCircleIcon,
  MailIcon,
  GithubIcon,
  GitlabIcon,
  GitPullRequestIcon,
  FacebookIcon,
  TwitterIcon,
  InstagramIcon,
  LinkedinIcon,
} from 'vue-feather-icons';
import Index from './components/Index.vue';

export default {
  name: 'App',
  components: {
    Index,
    SunIcon,
    MoonIcon,
    ArrowLeftCircleIcon,
    MailIcon,
    GithubIcon,
    GitlabIcon,
    GitPullRequestIcon,
    FacebookIcon,
    TwitterIcon,
    InstagramIcon,
    LinkedinIcon,
  },
  data: () => ({
    configLoaded: false,
    theme: '',
    metaTitle: '',
    title: '',
    description: '',
    defaultCurrency: '',
    acceptedCurrencies: [
      { value: 'usd', text: 'USD - United States Dollar' },
      { value: 'cad', text: 'CAD - Canadian Dollar' },
      { value: 'gbp', text: 'GBP - Pounds Sterling' },
      { value: 'aud', text: 'AUD - Australian Dollar' },
      { value: 'eur', text: 'EUR - Euro' },
    ],
    enableOneTime: true,
    enableMonthly: true,
    enableAnnual: true,
    backLink: '',
    emailLink: '',
    gitLink: '',
    githubLink: '',
    gitlabLink: '',
    facebookLink: '',
    twitterLink: '',
    instagramLink: '',
    linkedinLink: '',
  }),
  mounted() {
    const userPrefersDark = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
    let themePreference = 'dark';

    if (userPrefersDark) {
      themePreference = 'dark';
    } else {
      themePreference = 'light';
    }

    this.theme = localStorage.getItem('data-theme') || themePreference;

    this.axios.get('/api/site/config').then((response) => {
      // Passed to the Index component
      this.title = response.data.Title;
      this.description = response.data.Description;
      this.defaultCurrency = response.data.DefaultCurrency;
      // For this component
      this.metaTitle = response.data.MetaTitle;
      this.backLink = response.data.BackLink;
      this.emailLink = response.data.EmailLink;
      this.gitLink = response.data.GitLink;
      this.githubLink = response.data.GithubLink;
      this.gitlabLink = response.data.GitlabLink;
      this.facebookLink = response.data.FacebookLink;
      this.twitterLink = response.data.TwitterLink;
      this.instagramLink = response.data.InstagramLink;
      this.linkedinLink = response.data.LinkedinLink;
    }).then(() => {
      document.title = this.metaTitle;
      this.configLoaded = true;
    });
  },
  methods: {
    toggleTheme() {
      const newTheme = this.theme === 'light' ? 'dark' : 'light';
      this.theme = newTheme;
    },
  },
  watch: {
    theme() {
      const htmlElement = document.documentElement;
      if (this.theme === 'dark') {
        localStorage.setItem('data-theme', 'dark');
        htmlElement.setAttribute('data-theme', 'dark');
      } else {
        localStorage.setItem('data-theme', 'light');
        htmlElement.setAttribute('data-theme', 'light');
      }
    },
  },
};
</script>
<style lang="scss">
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}
body {
    background-color: var(--bg-color);
}
legend,label,h1,h2,h3,h4,h5,h6,p {
    color: var(--text-color);
}
.navIcons {
  color: var(--text-color);
}
.navIcons:hover {
  filter: invert(0.5);
}
#top-nav {
    border-top: 14px solid var(--top-nav-color);
}
.radio.btn-theme {
    background-color: #ffffff !important;
    border-color: var(--button-color) !important;
    color: var(--button-color) !important;
}
.radio.btn-theme.btn-sm.active {
    background-color: var(--button-color) !important;
    border-color: var(--button-color) !important;
    color: #ffffff !important;
}
.btn-theme {
    background-color: var(--button-color) !important;
    color: #ffffff !important;
}
.form-item {
  width: 100%;
  margin: auto;
}
</style>
