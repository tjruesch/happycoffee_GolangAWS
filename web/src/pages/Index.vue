<template>
    <app-layout>
        <div class="py-5 text-center">
            <h1 style="font-size: 3.5rem; color: #555;">happy&middot;<span style="color: #9e1b45;">coffee</span></h1>
            <span class="mb-5">Coffe <span style="color: #9e1b45;">&</span>more</span>
        </div>
        <h3 class="mb-3">Our Specialities</h3>
        <div>
            <b-table :data="data" :columns="columns"></b-table>
        </div>
        <div class="text-center mt-5">
            <a class="btn btn-outline-primary btn-lg text-primary" href="#">Download PDF Menue</a>
        </div>


    </app-layout>
</template>

<script>
    import AppLayout from '../layouts/App.vue'

    export default {
        components: {
            AppLayout,
        },
        data() {
            return {
                data: [
                    {drink: "Cold Brew", price: "4.50 EUR", happy_day: "Monday"},
                    {drink: "Cappuchino", price: "4.50 EUR", happy_day: "Wednesday"}
                ],
                columns: [
                {
                    field: 'drink',
                    label: 'Drink',
                },
                {
                    field: 'price',
                    label: 'Price',
                },
                {
                    field: 'happy_day',
                    label: 'Happy Day',
                }
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

    .old-date {
        text-decoration: line-through;
        opacity: 0.5;
    }
    .spacer {
        min-height: 100px;
    }
    .props {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
    .citation {
        font-size: 1.7rem;
    }
</style>
