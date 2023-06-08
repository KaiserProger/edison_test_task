<script setup lang="ts">

  import { useForm } from 'vee-validate';
  import { object, number } from 'yup';
  import { toTypedSchema } from '@vee-validate/yup';
  import { checkAnswers } from '@/network/checkExtrasenseAnswers';
  import router from '@/router';
  const { handleSubmit } = useForm({
    validationSchema: toTypedSchema(object().shape({
      number: number().min(10, "Too low!")
      .max(99, "Too much!").required("Please enter a number")
      .typeError("Must be a number").default(10),
    })),
  });

  const onSubmit = handleSubmit((values) => {
    checkAnswers(values.number);
    router.push("/");
  });
</script>

<template>
  <div class="row mt-4">
    <div class="col"></div>
    <div class="col xs-4">
      <h1>Enter two-digit number</h1>
      <form @submit="onSubmit">
        <div class="mb-3">
          <label for="numberInput" class="form-label">Enter a number</label>
          <input type="number" class="form-control" id="numberInput" aria-describedby="numberHelp">
        </div>
        <button type="submit" class="btn btn-primary">Submit</button>
      </form>
    </div>
    <div class="col"></div>
  </div>
</template>
