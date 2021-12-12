<template>
    <app-layout>
        <div class="hero w-100">
            <img src="../assets/media/hero.jpg" alt="Best Coffee in Town" width="100%">
            <span>&copy; Nathan Dumalo <a href="https://unsplash.com/photos/KixfBEdyp64">@ Unspash</a></span>
        </div>
        <div class="text-center">
            <h1 style="font-size: 3.5rem; color: #555;">happy&middot;<span style="color: #9e1b45;">coffee</span></h1>
            <span>Coffe <span style="color: #9e1b45;">&</span>more</span>
        </div>
        <h3 class="mb-3">Our Specialities</h3>
        <div>
            <b-table :data="products" :columns="columns"></b-table>
        </div>
        <div class="text-center mt-5">
            <a class="btn btn-outline-primary btn-lg text-primary" href="#">Download PDF Menue</a>
        </div>


    </app-layout>
</template>

<script>
    import AppLayout from '../layouts/App.vue'
    import axios from 'axios';

    export default {
        components: {
            AppLayout,
        },
        created() {
            axios.get("http://localhost:9090/v1/products")
                 .then(resp => {
                     resp.data.forEach(element => {
                         console.log(element)
                         if (element['discount'] !== 0) {
                             let priceOld = element['price']
                             let priceNew = (parseFloat(element['price']) * (100-element['discount'])/100 ).toPrecision(3)
                             element['price'] = `
                                <span style="text-decoration: line-through; opacity: 0.5;">${priceOld}</span>
                                <span class="pl-2" style="font-weight: bold;"><em>${priceNew} EUR</em></span>
                                `
                         }
                     });
                     this.products = resp.data
                 })
                 .catch(function (error) {
                   console.log(error)
                 })
        },
        data() {
            return {
                products: [],
                columns: [
                {
                    field: 'name',
                    label: 'Drink',
                },
                {
                    field: 'price',
                    label: 'Price',
                },
                {
                    field: 'happy_day',
                    label: 'Happy Day',
                },
                ]
            }
        }
    }
</script>

<style scoped>
    h3 {
        margin: 40px 0 0;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        display: inline-block;
        margin: 0 10px;
    }

    a {
        color: #42b983;
    }   
    
    a.btn {
        color: #fff 
    }
    body {
        background-color: #f5f5f5;
    }
    .old-price {
        text-decoration: line-through;
    }
</style>
