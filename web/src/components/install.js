import ACE from './Ace'
import ProblemModal from "./ProblemModal";
import Back from "./Back";

export default {
  install: function (Vue) {
    Vue.component(ACE.name, ACE);
    Vue.component(ProblemModal.name, ProblemModal);
    Vue.component(Back.name, Back);
  }
}
