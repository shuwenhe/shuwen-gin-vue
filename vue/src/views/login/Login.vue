<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col md="8"
      offset="2"
      lg="6"
      offset-lg="3"
      >
    <b-card title="Login">
    <b-form>
      <b-form-group label="Phone">
        <b-form-input
          v-model="$v.user.phone.$model"
          type="number"
          placeholder="Input your phone"
          :state="validateState('phone')"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState('phone')">
          Mobile phone number does not meet the requirements!
        </b-form-invalid-feedback>
      </b-form-group>
      <b-form-group label="Password">
        <b-form-input
          v-model="$v.user.password.$model"
          type="password"
          placeholder="Input your password"
        ></b-form-input>
        <b-form-invalid-feedback :state="validateState('password')">
          Password must be greater than or equal to 6 digits!
        </b-form-invalid-feedback>
        </b-form-group>
        <b-form-group>
        <b-button
        @click="login"
        variant="outline-primary" block>Login</b-button>
        </b-form-group>
    </b-form>
    </b-card>
    </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength } from 'vuelidate/lib/validators'
import customValidator from "@/helper/validator"

export default {
  data() {
   return {
     user: {
      phone: "",
      password: "",
    },
    validation: null,
   };
  },
    validations: {
    user: {
      phone: {
      required,
      phone: customValidator.phoneValidator,
      },
    password: {
      required,
      minLength: minLength(6),
    },
    },
  },
  methods: {
    validateState(name) {
      const { $dirty, $error } = this.$v.user[name] // es6解构赋值
      return $dirty ? !$error : null
    },
    login() {
      console.log("register")
    },
  },
}
</script>

<style lang="scss" scoped>

</style>
