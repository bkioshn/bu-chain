import { Client, registry, MissingWalletError } from 'bu-chain-client-ts'

import { Params } from "bu-chain-client-ts/buchain.exchange/types"
import { ExchangeRate } from "bu-chain-client-ts/buchain.exchange/types"


export { Params, ExchangeRate };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				ExchangeRate: {},
				ExchangeRateAll: {},
				ExchangeAmount: {},
				ExchangePairs: {},
				
				_Structure: {
						Params: getStructure(Params.fromPartial({})),
						ExchangeRate: getStructure(ExchangeRate.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getExchangeRate: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ExchangeRate[JSON.stringify(params)] ?? {}
		},
				getExchangeRateAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ExchangeRateAll[JSON.stringify(params)] ?? {}
		},
				getExchangeAmount: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ExchangeAmount[JSON.stringify(params)] ?? {}
		},
				getExchangePairs: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ExchangePairs[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: buchain.exchange initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.BuchainExchange.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryExchangeRate({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.BuchainExchange.query.queryExchangeRate( key.index)).data
				
					
				commit('QUERY', { query: 'ExchangeRate', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryExchangeRate', payload: { options: { all }, params: {...key},query }})
				return getters['getExchangeRate']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryExchangeRate API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryExchangeRateAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.BuchainExchange.query.queryExchangeRateAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.BuchainExchange.query.queryExchangeRateAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ExchangeRateAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryExchangeRateAll', payload: { options: { all }, params: {...key},query }})
				return getters['getExchangeRateAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryExchangeRateAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryExchangeAmount({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.BuchainExchange.query.queryExchangeAmount( key.denom,  key.amount,  key.exchangeToken)).data
				
					
				commit('QUERY', { query: 'ExchangeAmount', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryExchangeAmount', payload: { options: { all }, params: {...key},query }})
				return getters['getExchangeAmount']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryExchangeAmount API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryExchangePairs({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.BuchainExchange.query.queryExchangePairs(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.BuchainExchange.query.queryExchangePairs({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ExchangePairs', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryExchangePairs', payload: { options: { all }, params: {...key},query }})
				return getters['getExchangePairs']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryExchangePairs API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateExchangeRate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.BuchainExchange.tx.sendMsgCreateExchangeRate({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateExchangeRate:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateExchangeRate:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateExchangeRate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.BuchainExchange.tx.sendMsgUpdateExchangeRate({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateExchangeRate:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateExchangeRate:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgExchangeToken({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.BuchainExchange.tx.sendMsgExchangeToken({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgExchangeToken:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgExchangeToken:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteExchangeRate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.BuchainExchange.tx.sendMsgDeleteExchangeRate({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteExchangeRate:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteExchangeRate:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateExchangeRate({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.BuchainExchange.tx.msgCreateExchangeRate({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateExchangeRate:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateExchangeRate:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateExchangeRate({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.BuchainExchange.tx.msgUpdateExchangeRate({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateExchangeRate:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateExchangeRate:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgExchangeToken({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.BuchainExchange.tx.msgExchangeToken({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgExchangeToken:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgExchangeToken:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteExchangeRate({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.BuchainExchange.tx.msgDeleteExchangeRate({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteExchangeRate:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteExchangeRate:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
