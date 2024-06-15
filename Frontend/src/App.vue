<template>
  <Header />
  <div>
    <Body :Emails="emails"/>
  </div>
  <Footer />
</template>

<script >
import Body  from "../src/components/Body.vue"
import Header from "../src/components/Header.vue"
import Footer from "../src/components/Footer.vue"

export default {
  name: 'App',
  components: {
    Header,
    Footer,
    Body
  },
  data() {
    return {
      emails: [
        {
          subject: "changes made on 12/04 to 11/30",
          from: "valarie.sabo@enron.com",
          to: "stacey.white@enron.com",
          body: `The gas book liquidated twice, so our QTD and YTD was understated...`
        },
        {
          subject: "Enron Mentions - 03-04-01",
          from: "john.arnold@enron.com",
          to: "greg.whalley@enron.com",
          body: `In numbers, LT books QTD and YTD increased on 11/30 by approximately...`
        }
      ]
    }
  }
}
</script>

<style scoped>
.container {
  max-width: 100%;
}
</style>

