function (doc) {
  if (doc.result !== 'SUCCESS' && doc.stages && Array.isArray(doc.stages)) {
      emit(doc.id, { result: doc.result, lastStage: doc.stages[doc.stages.length-2].name, timeReadable: doc.timeReadable});
  }
}
