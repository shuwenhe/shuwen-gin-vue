<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col md="8"
      offset="2"
      lg="6"
      offset-lg="3"
      >
    <b-card title="Register">
    <b-form>
      <b-form-group label="Name">
        <b-form-input
          v-model="user.name"
          type="text"
          placeholder="Input your name"
        ></b-form-input>
      </b-form-group>
      <b-form-group label="Phone">
        <b-form-input
          v-model="user.phone"
          type="number"
          placeholder="Input your phone"
        ></b-form-input>
         <b-form-text
         id="password-help-block"
         text-variant="danger"
         v-if="showPhone"
         >
           phone num must 11 digits!
         </b-form-text>
      </b-form-group>
      <b-form-group label="Password">
        <b-form-input
          v-model="user.password"
          type="password"
          placeholder="Input your password"
        ></b-form-input>
        </b-form-group>
        <b-form-group>
        <b-button
        @click="register"
        variant="outline-primary" block>Register</b-button>
        </b-form-group>
    </b-form>
    </b-card>
    </b-col>
    </b-row>
  </div>
</template>

<script>
export default {
  data() {
   return {
     user: {
      name: "",
      phone: "",
      password: "",
    },
    showPhone: false,
   }
  },
  methods: {
    register() {
      if (user.phone.length != 11) {
        this.showPhone = true
        return
      }
      console.log("register")
    },
  },
}
</script>

<style lang="scss" scoped>

</style>
