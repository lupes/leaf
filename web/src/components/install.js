import ACE from './Ace'
import ProblemModal from "./ProblemModal";

export default {
  install: function (Vue) {
    Vue.component(ACE.name, ACE);
    Vue.component(ProblemModal.name, ProblemModal);
  }
}
