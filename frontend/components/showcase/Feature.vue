<template>
  <div >
    <div class="has-tooltip">
      <span class='tooltip text-sm absolute rounded shadow-lg bg-accent p-1 my-2 mx-8 z-30'>{{ feature.Description }}</span>
      <button class="p-2 text-sm relative flex" @click="showSettings = !showSettings">
        <span >{{feature.Name}}</span>
        <svg v-if="feature.Settings.length !== 0" class="w-5 h-5" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M12 15.713L18.01 9.70299L16.597 8.28799L12 12.888L7.40399 8.28799L5.98999 9.70199L12 15.713Z" fill="currentColor"></path>
        </svg>
      </button>
    </div>
    <ul v-show="feature.Settings.length !== 0 && showSettings" class="p-2 text-sm">
      <li v-for="s in feature.Settings" class="flex relative border-l border-t flex-wrap">
        <div class="has-tooltip">
          <span v-if="s.description.length > 1" class='tooltip text-sm absolute rounded shadow-lg bg-accent p-1 -my-16 mx-8 z-50'>{{ s.description }}</span>
          <span>Name: {{s.name}}</span>

          <span>Value: {{s.value}}</span>
          <div v-if="s.type === 'number'">
            <a>Min: {{s.min}}</a>
            <a>Max: {{s.max}}</a>
          </div>
          <span v-if="s.type === 'Enum'">
            Values:
            <ul>
              <li v-for="vs in s.values">
                <a>{{vs}}</a>
              </li>
            </ul>
          </span>
        </div>
        <ul v-show="s.subsettings && s.subsettings.length !== 0" class="p-2 text-sm mx-1">
          SubSettings:
          <li v-for="ss in s.subsettings" class="flex flex-col relative border-l border-t">
            <div class="has-tooltip">
              <span v-if="ss.description.length > 1" class='tooltip text-sm absolute rounded shadow-lg bg-accent p-1 -my-16 mx-8 z-50'>{{ ss.description }}</span>
              <span class="">Name: {{ss.name}}</span>
              <span>Value: {{ss.value}}</span>
              <div v-if="ss.type === 'number'">
                <a>Min: {{ss.min}}</a>
                <a>Max: {{ss.max}}</a>
              </div>
              <span v-if="ss.type === 'Enum'">
                Values:
                <ul>
                  <li v-for="vs in ss.values">
                    <a>{{vs}}</a>
                  </li>
                </ul>
              </span>
            </div>
          </li>
        </ul>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

interface Feature {
  Name: string;
  Description: string;
  Settings: Setting[];
}

interface Setting {
  name: string;
  description: string;
  type: string;
  value: any;
  min?: number;
  max?: number;
  values: string[];
  subsettings?: Setting[]
}
export default Vue.extend({
  name: "Feature",
  props: {
    fe: {
      type: Object as () => Feature
    }
  },
  data() {
    return {
      feature: this.fe as Feature,
      showSettings: false
    }
  }
})
</script>

<style scoped>

</style>
